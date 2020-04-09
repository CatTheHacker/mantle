"use strict";
//
import { setDataBinding, deActivateChild } from "./../util.js";
import { output } from "./../ui.util.js";
import * as ui from "./../ui.js";
import * as api from "./../api/index.js";

//
customElements.define("x-message", class extends HTMLElement {
    constructor() {
        super();
    }
    get time() {
        return new moment(this.querySelector(".ts").getAttribute("title"), moment.defaultFormat);
    }
    connectedCallback() {
        this._uid = this.getAttribute("uuid");
        this._author = this.getAttribute("author");
        //
        this.addEventListener("click", (e) => {
            /** @type {Element[]} */
            const fl = e.composedPath().filter((v) => v instanceof Element && v.matches("[uuid]"));
            if (fl.length === 0) return;
            const et = fl[0];
            if (e.ctrlKey) {
                et.classList.toggle("selected");
            }
        });
        this.querySelector(".usr").addEventListener("click", async (e) => {
            const userN = await api.M.users.get(e.target.parentElement._author);
            setDataBinding("pp_user_name", userN.name);
            setDataBinding("pp_user_nickname", userN.nickname);
            setDataBinding("pp_user_id", userN.id);
            setDataBinding("pp_user_uuid", userN.uuid);
            setDataBinding("pp_user_provider", userN.provider);
            setDataBinding("pp_user_snowflake", userN.snowflake);
            const pp = document.querySelector("dialog.popup.user");
            const ppr = pp.querySelector("ol");
            deActivateChild(ppr);
            const pps = pp.querySelector("div ol");
            deActivateChild(pps);
            const rls = await userN.getRoles();
            for (const item of rls) {
                const ppra = ppr.querySelector(`[data-role="${item.uuid}"]`);
                if (ppra === null) continue;
                ppra.classList.add("active");
                const ppsa = pps.querySelector(`[data-role="${item.uuid}"]`);
                if (ppsa === null) continue;
                ppsa.classList.add("active");
            }
            pp.setAttribute("open","");
            pp.style.top = e.y+"px";
            pp.style.left = e.x+"px";
        });
    }
});

//
document.addEventListener("keydown", async (e) => {
    if (e.key !== "Delete") return;
    if (document.querySelector("body > [open]") !== null) return;
    if (document.activeElement !== document.body) return;
    const sel = output.getChannel(output.active_channel_uid).selected();
    if (sel.length === 0) return;
    await Swal.fire({
        title: "Are you sure you want to delete?",
        text: "You won't be able to revert this!",
        type: "warning",
        showCancelButton: true,
    })
    .then(async (r) => {
        if (!r.value) return;
        const m2d = sel.filter((v) => v._author === ui.volatile.me.uuid).map((v) => v._uid);
        await api.M.channels.with(output.active_channel_uid).messages.delete(m2d);
    });
});