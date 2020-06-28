<template>
    <header class="sticky row">
            <a href="#" class="logo"><router-link to="/">ÖPNV</router-link></a>
            <router-link class="button" to="/schedule">Fahrpläne</router-link>
            <router-link class="button" to="/ticket">Tickets</router-link>
            <a href="#" class="button hidden-sm" style="color: lightgrey">Carsharing</a>
            <a href="#" class="button hidden-sm" style="color: lightgrey">City Roller</a>
        <div v-if="user == false">
            <router-link class="button" id="login" to="/login">
                Login
            </router-link>
        </div>
        <div v-if="user == true">
            <router-link class="button" id="login" to="/myaccount">
                Mein Konto
            </router-link>
            <a href="#" class="button" @click.prevent="logoutUser()" id="logout">
                Logout
            </a>
        </div>
    </header>
</template>
<script>
import firebase from 'firebase';

export default {
    data() {
        return {
            user: '',
            email: '',
        };
    },
    created() {
        firebase.auth().onAuthStateChanged((user) => {
            if (user) {
                this.user = true;
                this.email = user.email;
            } else {
                this.user = false;
            }
        });
    },
    methods: {
        logoutUser() {
            firebase.auth().signOut().then(() => {
                // Logout erfolgreich
                alert('Logout Erfoglreich!');
            }, (error) => {
                alert('Logout fehlgeschlagen');
            });
        },
    },
};
</script>
