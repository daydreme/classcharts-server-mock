const loginForm = document.getElementById("login")
const schoolForm = document.getElementById("school")
if(loginForm)
    loginForm.addEventListener("submit", async (event) => {
        event.preventDefault()

        logOut()

        const res = await apiv2student("login", new FormData(loginForm))
        if(res.meta?.session_id) localStorage.setItem("token", res.meta?.session_id)

        await doPing()
    })

if(schoolForm)
    schoolForm.addEventListener('submit', (event) => {
        event.preventDefault()

        localStorage.setItem("sid", new FormData(schoolForm).get("sId"))
        checkSchoolId()
    })

async function apiv2student(endpoint, data, authed=true, action="POST") {
    let headers = {}
    authed && localStorage.getItem("token") ? headers['Authorization'] = `Basic ${localStorage.getItem("token")}` : ""

    const res = await (await fetch(
        "/apiv2student/" + endpoint,
        {
            headers,
            body: data,
            method: action
        }
    )).json()

    if(res.success !== 1) {
        alert(res.error ?? "An unknown error occured.")
        return {}
    }

    return res
}

async function doPing() {
    const token = localStorage.getItem("token")
    if(token) {
        const res = await apiv2student("ping", null)
        document.getElementById("userName").innerText += `${res.data.user.name} (id ${res.data.user.id})`

    }
}

async function logOut() {
    localStorage.clear()
    document.getElementById("userName").innerText = "No one."
}

function checkReplacements() {
    document.querySelectorAll("*[sid]").forEach(siddable => {
        siddable.innerText += `School ID: ${localStorage.getItem("sid") ?? "None set."}`
    })
}

(async () => {
    await doPing()
    checkReplacements()
})()