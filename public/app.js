import { API } from "./services/API.js";
import Router from "./services/Router.js";

import "./components/HomePage.js";
import "./components/AnimatedLoading.js";

window.addEventListener("DOMContentLoaded", () => {
    app.Router.init();
});

window.app = {
    search: (event) => {
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
    },

    api: API,
    Router,
};
