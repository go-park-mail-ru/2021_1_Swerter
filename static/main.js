var app = document.getElementById("app");

headerContent = {
    profile: {
        link: '/profile',
    },
    news: {
        link: '/news '
    },
    auth: {
        link: '/auth',
    },
}

function home() {
    const headerContent = document.createElement('div');
    headerContent.classList.add('info-user');

    const centralGradient = document.createElement('div');
    centralGradient.classList.add('about-user');
    ava = document.createElement('div');
    ava.className = 'ava';
    centralGradient.appendChild(ava);

    headerContent.appendChild(centralGradient);

    const bottomContent = document.createElement('div');
    bottomContent.classList.add('content-user');

    const post = document.createElement('div');
    post.className = 'post';
    bottomContent.appendChild(post);
    getContent(bottomContent);
    app.appendChild(headerContent);
    app.appendChild(bottomContent);
}

function getContent(bottomContent) {
    fetch('https://jsonplaceholder.typicode.com/posts')
        .then(response => response.json())
        .then(json => {
            json
                .map(post => {
                    const newPost = document.createElement('div');
                    newPost.className = 'post';
                    newPost.innerHTML = post['body'];
                    return newPost;
                })
                .forEach(newPost => bottomContent.appendChild(newPost))
        })

}

function news() {
}

function generateMaket() {
    generateHeader();
    //generateBody();
    //generateBottom();
}

function changePage(link) {
    app.innerHTML = "";
    generateMaket();
    if (link == '/profile') {
        home();
    } else if (link == '/news') {
        news();
    }
}

function generateHeader() {
    header = document.createElement('div');
    Object
        .keys(headerContent)
        .map((item) => {
            nav = document.createElement('div');
            nav.classList.add('top-link');
            nav.innerHTML = item[0].toUpperCase() + item.slice(1);
            nav.dataset.link = headerContent[item]['link'];
            nav.addEventListener('click', (evt) => {
                const {target} = evt;
                // console.log(target.dataset.link);
                changePage(target.dataset.link);
            });
            return nav;
        })
        .forEach((nav) => {
            header.appendChild(nav);
        })
    header.classList.add('header');
    app.appendChild(header);
}

changePage('/profile');