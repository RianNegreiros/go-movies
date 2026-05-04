import { API } from "./services/API.js";

import "./components/HomePage.js";
import "./components/AnimatedLoading.js";

window.app = {
    search: (event) => {
        event.preventDefault();
        const q = document.querySelector("input[type=search]").value;
    },

    api: API,
};
