function toggleMenu() {
    document
        .getElementById("accountMenu")
        .classList
        .toggle("show");
}

window.onclick = function(e) {
    if (!e.target.closest(".account")) {
        document
            .getElementById("accountMenu")
            .classList
            .remove("show");
    }
};


async function api(url, options = {}) {

    const request = () => fetch(url, {
        ...options,
        credentials: "include",
        headers: {
            ...(options.headers || {}),
            Authorization:
                "Bearer " + sessionStorage.getItem("access_token")
        }
    });


    let response = await request();


    if (response.status === 401) {

        const ok = await refresh();

        if (!ok) {
            return response;
        }

        response = await request();
    }


    return response;
}


async function refresh() {

    const response = await fetch("/api/auth/refresh", {
        method: "POST",
        credentials: "include"
    });


    if (!response.ok) {
        logoutLocal();
        return false;
    }


    const data = await response.json();

    sessionStorage.setItem(
        "access_token",
        data.access_token
    );


    return true;
}


async function loadUser() {

    const response = await api("/api/users/me");


    if (!response.ok) {
        return;
    }


    const user = await response.json();


    document.querySelector(".user-name").textContent =
        user.name;

    document.querySelector(".user-email").textContent =
        user.email;

    document.querySelector(".user-created").textContent =
        new Date(user.created_at)
            .toLocaleDateString();
}



async function logout() {

    await fetch("/api/auth/logout", {
        method: "POST",
        credentials: "include"
    });


    logoutLocal();
}



async function deleteAccount() {

    const response = await api("/api/users/me", {
        method: "DELETE"
    });


    if (!response.ok) {

        const error = await response.text();

        console.error(error);

        alert("Failed to delete account");

        return;
    }


    logoutLocal();
}



function logoutLocal() {

    sessionStorage.removeItem(
        "access_token"
    );

    window.location.href = "/";
}



async function initStorage() {

    const ok = await refresh();

    if (!ok) {
        return;
    }


    await loadUser();
}



document.addEventListener(
    "DOMContentLoaded",
    initStorage
);
