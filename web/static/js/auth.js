async function parseJSON(response) {
    const text = await response.text();
    try {
        return JSON.parse(text);
    }
    catch {
        return text;
    }

}

async function signup() {
    const response = await fetch("/auth/signup", {
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
    await loginUser(
        document.getElementById("signup-email").value,
        document.getElementById("signup-password").value
    );
}

async function login() {
    await loginUser(
        document.getElementById("login-email").value,
        document.getElementById("login-password").value
    );
}



async function loginUser(email,password) {
    const response = await fetch("/auth/login", {
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
        return;
    }
    sessionStorage.setItem(
        "access_token",
        data.access_token
    );
    window.location.href="/storage";
}
