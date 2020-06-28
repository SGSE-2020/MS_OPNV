<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <div v-if="this.user == true" id="login">
                        Diesen Bereich sieht man nur wenn man eingeloggt ist.
                    </div>
                    <div v-if="this.user == false">
                        Um Tickets zu kaufen müssen sie sich einloggen!
                    </div>
                    <!--<button class="primary" @click="testBackendCall()">TestBackendCall</button>
                    <p v-if="this.text != ''">{{this.text}}</p>
                    <button class="primary" @click="testGetUsersCall()">TestGetUserCall</button>
                    <ul v-if="this.users.length != 0">
                        <li v-for="(user, i) in this.users" :key="i">
                            {{ user.token }}
                        </li>
                    </ul>-->
                </div>
                <TheSidebar />
            </div>
        </div>
    </div>
</template>

<script>
import firebase from 'firebase';
import axios from 'axios';
import TheHeader from '../components/TheHeader.vue';
import TheSidebar from '../components/TheSidebar.vue';

export default {
    name: 'Ticket',
    components: {
        TheHeader,
        TheSidebar,
    },
    data() {
        return {
            text: '',
            user: '',
            users: [],
        };
    },
    created() {
        firebase.auth().onAuthStateChanged((user) => {
            if (user) {
                this.user = true;
            } else {
                this.user = false;
            }
        });
    },
    methods: {
        testBackendCall() {
            axios.get(`${process.env.VUE_APP_BACKEND_HOST}/api`)
                .then((response) => {
                    console.log(response.data);
                    this.text = response.data;
                })
                .catch((e) => {
                this.errors.push(e);
                });
        },
        testGetUsersCall() {
            axios.get(`${process.env.VUE_APP_BACKEND_HOST}/users`)
                .then((response) => {
                    this.users = response.data;
                })
                .catch((e) => {
                    this.errors.push(e);
                });
        },
    },
};
</script>
