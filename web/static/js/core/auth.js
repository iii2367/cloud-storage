async function login() {
    
    const logged = await loginUser(
        document.getElementById("login-email").value,
        document.getElementById("login-password").value
    );

    if (!logged) {
        return;
    }
    window.location.href="/storage";
}

async function loginUser(email,password) {

    const response = await fetch("/api/auth/login", {
        method:"POST",
        credentials:"include",
        headers: {
            "Content-Type":"application/json"
        },
        body:JSON.stringify( {
            email,
            password
        })
    });
    const data = await parseJSON(response);
    if (!response.ok) {
        alert(data.error || "Login failed");
        return false;
    }
    sessionStorage.setItem(
        "access_token",
        data.access_token
    );
    return true;
}


async function logout() {

    await fetch("/api/auth/logout", {
        method: "POST",
        credentials: "include"
    });

    logoutLocal();
}

function logoutLocal() {

    sessionStorage.removeItem(
        "access_token"
    );

    window.location.href = "/";
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

async function createRootFolder() {

    const response = await fetch("/api/storage/root", {
        method: "POST",
        credentials: "include",
        headers: {
            "Authorization":
                "Bearer " + sessionStorage.getItem("access_token")
        }
    });
    if (!response.ok) {
        const data = await parseJSON(response);
        console.error(data);
        return false;
    }
    return true;
}

async function signup() {
 
    const response = await fetch("/api/auth/signup", {
        method:"POST",
        headers: {
            "Content-Type":"application/json"
        },
        credentials:"include",
        body:JSON.stringify( {
            name:
            document.getElementById("signup-name").value,
            email:
            document.getElementById("signup-email").value,
            password:
            document.getElementById("signup-password").value
        })
    });
    if (!response.ok) {
        const data = await parseJSON(response);
        alert(data.error || "Signup error");
        return;
    }
    const logged = await loginUser(
        document.getElementById("signup-email").value,
        document.getElementById("signup-password").value
    );
    
    if (!logged) {
        return;
    }

    const folroot = await createRootFolder();

    if (!folroot) {
        return;
    }

    window.location.href="/storage";
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
