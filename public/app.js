import { API } from "./services/API.js";
import Router from "./services/Router.js";

import "./components/HomePage.js";
import "./components/AnimatedLoading.js";

window.addEventListener("DOMContentLoaded", () => {
    app.Router.init();
});

window.app = {
    api: API,
    Router,
    search: (event) => {
        event.preventDefault();
        const keywords = document.querySelector("input[type=search]").value;
        if (keywords.length > 1) {
            app.Router.go(`/movies?q=${keywords}`);
        }
    },
    showError: (message = "There was an error", goToHomePage = true) => {
        document.getElementById("alert-modal").showModal();
        document.querySelector("#alert-modal p").textContent = message;
        if (goToHomePage) app.Router.go("/");
    },
    closeError: () => {
        document.getElementById("alert-modal").close();
    },
    searchOrderChange: (order) => {
        const urlParams = new URLSearchParams(window.location.search);
        const q = urlParams.get("q");
        const genre = urlParams.get("genre") ?? "";
        app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
    },
    searchFilterChange: (genre) => {
        const urlParams = new URLSearchParams(window.location.search);
        const q = urlParams.get("q");
        const order = urlParams.get("order") ?? "";
        app.Router.go(`/movies?q=${q}&order=${order}&genre=${genre}`);
    },
};
