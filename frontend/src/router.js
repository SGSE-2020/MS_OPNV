// import firebase from 'firebase';
import Vue from 'vue';
import Router from 'vue-router';

import Home from './views/Home.vue';
import Login from './views/Login.vue';
import AddUser from './views/AddUser.vue';

Vue.use(Router);

const router = new Router({
    routes: [{
            path: '*',
            redirect: '/',
        },
        {
            path: '/',
            redirect: '/home',
        },
        {
            path: '/login',
            name: 'Login',
            component: Login,
        },
        {
            path: '/home',
            name: 'home',
            component: Home,
            // meta: {
            //     requiresAuth: true,
            // },
        },
        {
            path: '/user',
            name: 'User',
            component: AddUser,
        },
    ],
});

// router.beforeEach((to, from, next) => {
//     const { currentUser } = firebase.auth().currentUser;
//     const requiresAuth = to.matched.some((record) => (record.meta.requiresAuth));

//     if (requiresAuth && !currentUser) next('login');
//     else if (!requiresAuth && currentUser) next('home');
//     else next();
// });

export default router;
