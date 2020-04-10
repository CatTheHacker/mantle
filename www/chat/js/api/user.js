"use strict";
//
import * as api from "./index.js";

//
/** @type {Map<string,User} */
export const cache = new Map();

//
export class User {
    constructor(o) {
        if (o === null) {
            this.is_null = true;
            return;
        }
        Object.assign(this, o);
        this.is_null = false;
        cache.set(this.uuid, this);
    }
    /** @returns {Promise<api.Role[]>} */
    getRoles() {
        return Promise.all(this.roles.map((v) => api.M.roles.get(v))).then((l) => {
            return l.sort((a,b) => a.position - b.position);
        });
    }
    getName() {
        if (this.nickname.length > 0) {
            return this.nickname;
        }
        return this.name;
    }
}
