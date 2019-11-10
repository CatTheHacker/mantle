/**
 * @see https://github.com/nektro/mantle
 * @author Meghan Denny <https://nektro.net>
 */
//
import * as ut from "./util.js";

//
(async function() {
    //
    await fetch("/api/about").then(x => x.json()).then(x => {
        ut.el_2.innerText = x.name;
    });

    //
    await fetch("/api/users/@me").then(x => x.json()).then(x => {
        if (x.success === false) {
            location.assign("../");
        }
        else {
            const u = x.message.me;
            const n = u.nickname || u.name;
            ut.el_3.children[0].textContent = `@${n}`;

            const p = x.message.perms;
            for (const key in p) {
                document.querySelectorAll(`[data-requires^="${key}"]`).forEach((el) => {
                    el.removeAttribute("hidden");
                });
            }
        }
    });
})();
