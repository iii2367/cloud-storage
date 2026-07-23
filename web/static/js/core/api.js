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

async function parseJSON(response) {

    const text = await response.text();
    try {
        return JSON.parse(text);
    }
    catch {
        return text;
    }
}
