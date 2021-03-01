const application = document.getElementById('app');

const config = {
    menu: {
        href: '/menu',
        text: 'Меню!',
        open: menuPage,
    },
    signup: {
        href: '/signup',
        text: 'Зарегистрироваться!',
        open: signupPage,
    },
    login: {
        href: '/login',
        text: 'Авторизоваться!',
        open: loginPage,
    },
    profile: {
        href: '/profile',
        text: 'Профиль',
        open: profilePage,
    },
    about: {
        href: '/about',
        text: 'Контакты',
    }
}

function createInput(type, text, name) {
    const input = document.createElement('input');
    input.type = type;
    input.name = name;
    input.placeholder = text;

    return input;
}

function menuPage() {
    application.innerHTML = '';

    Object
        .entries(config)
        .map(([menuKey, {text, href}]) => {
            const menuItem = document.createElement('a');
            menuItem.href = href;
            menuItem.textContent = text;
            menuItem.dataset.section = menuKey;

            return menuItem;
        })
        .forEach(element => application.appendChild(element))
    ;
}

function signupPage() {
    application.innerHTML = '<h1>Регистрация!</h1>';

    const form = document.createElement('form');

    const loginInput = createInput('login', 'Емайл', 'login');
    const passwordInput = createInput('password', 'Пароль', 'password');

    const submitBtn = document.createElement('input');
    submitBtn.type = 'submit';
    submitBtn.value = 'Зарегистрироваться!';

    const back = document.createElement('a');
    back.href = '/menu';
    back.textContent = 'Назад';
    back.dataset.section = 'menu';

    form.appendChild(loginInput);
    form.appendChild(passwordInput);
    form.appendChild(submitBtn);
    form.appendChild(back);

    application.appendChild(form);

    form.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const login = loginInput.value.trim();
        const password = passwordInput.value.trim();

        ajax('POST', '/register', {login, password}, (status, responseBody) => {
            console.log(status)
            console.log(responseBody)
        });
    });
}

function loginPage() {
    application.innerHTML = '';
    const form = document.createElement('form');

    const loginInput = createInput('login', 'Емайл', 'login');
    const passwordInput = createInput('password', 'Пароль', 'password');

    const submitBtn = document.createElement('input');
    submitBtn.type = 'submit';
    submitBtn.value = 'Авторизироваться!';

    const back = document.createElement('a');
    back.href = '/menu';
    back.textContent = 'Назад';
    back.dataset.section = 'menu';

    form.appendChild(loginInput);
    form.appendChild(passwordInput);
    form.appendChild(submitBtn);
    form.appendChild(back);


    form.addEventListener('submit', (evt) => {
        evt.preventDefault();

        const login = loginInput.value.trim();
        const password = passwordInput.value.trim();

        ajax(
            'POST',
            '/login',
            {login, password},
            (status, response) => {
                if (status === 200) {
                    // profilePage();
                    console.log(JSON.parse(response))
                } else {
                    const {error} = JSON.parse(response);
                    alert(error);
                }
            }
        )

    });

    application.appendChild(form);
}

function profilePage() {
    application.innerHTML = '';

    ajax('GET', '/profile', null, (status, responseText) => {
        let isAuthorized = false;

        if (status === 200) {
            isAuthorized = true;
        }

        if (status === 401) {
            isAuthorized = false;
        }


        if (isAuthorized) {
            const responseBody = JSON.parse(responseText);

            console.log(responseBody)

            const logout = document.createElement('a')
            logout.href = '#'
            logout.innerText = 'Logout'
            logout.dataset.section = 'menu';
            application.appendChild(logout)
            logout.addEventListener('click', e => {
                const {target} = e;
                e.preventDefault()
                ajax('POST', '/logout', null, (status, responseText) => {
                    //handle if POST failed
                    if (status == 200) {
                        console.log("logout")
                    }
                })

            })


            const back = document.createElement('a');
            back.href = '/menu';
            back.textContent = 'Назад';
            back.dataset.section = 'menu';

            application.appendChild(back);

            return;
        }

        alert('АХТУНГ! НЕТ АВТОРИЗАЦИИ');

        loginPage();
    });
}

menuPage();

application.addEventListener('click', e => {
    const {target} = e;

    if (target instanceof HTMLAnchorElement) {
        e.preventDefault();
        config[target.dataset.section].open();
    }
});

function ajax(method, url, body = null, callback) {
    const xhr = new XMLHttpRequest();
    xhr.open(method, url, true);
    xhr.withCredentials = true;

    xhr.addEventListener('readystatechange', function () {
        if (xhr.readyState !== XMLHttpRequest.DONE) return;

        callback(xhr.status, xhr.responseText);
    });

    if (body) {
        xhr.setRequestHeader('Content-type', 'application/json; charset=utf8');
        xhr.send(JSON.stringify(body));
        return;
    }

    xhr.send();
}
