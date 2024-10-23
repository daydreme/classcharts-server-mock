async function loadAnnouncements() {
    const res = await (await apiv2student('announcements/1', null, action = 'GET')).json()
    for (let announcement of res.data) {
        document.getElementById("announcements").innerHTML += `<div class="announcement">
            <h2>${announcement.title}</h2>
            <p>${announcement.content}</p>
        </div>`
    }
}