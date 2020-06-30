// import firebase from 'firebase';
import Vue from 'vue';
import Router from 'vue-router';

import Home from './views/Home.vue';
import Login from './views/Login.vue';
import Schedule from './views/Schedule.vue';
import Ticket from './views/Ticket.vue';
import MyAccount from './views/MyAccount.vue';
import Parkspace from './views/Parkspace.vue';

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
            name: 'Home',
            component: Home,
            // meta: {
            //     requiresAuth: true,
            // },
        },
        {
            path: '/schedule',
            name: 'Schedule',
            component: Schedule,
        },
        {
            path: '/ticket',
            name: 'Ticket',
            component: Ticket,
        },
        {
            path: '/myaccount',
            name: 'MyAccount',
            component: MyAccount,
        },
        {
            path: '/parkspace',
            name: 'Parkspace',
            component: Parkspace,
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
