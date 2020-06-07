<template>
    <header class="sticky row">
        <div class="col-md-2">
            <a href="#" class="logo"><router-link to="/">ÖPNV</router-link></a>
        </div>
        <div class="col-md-7" >
            <router-link class="button hidden-sm" to="/schedule">Fahrpläne</router-link>
            <router-link class="button hidden-sm" to="/ticket">Tickets</router-link>
            <a href="#" class="button hidden-sm" style="color: lightgrey">Carsharing</a>
            <a href="#" class="button hidden-sm" style="color: lightgrey">City Roller</a>
        </div>
        <div v-if="user == false" class="col-md">
            <router-link class="button hidden-sm" id="login" to="/login">
            Login
            </router-link>
        </div>
        <div v-if="user == true" class="col-md">
            <router-link class="button hidden-sm" id="login" to="/login">
            <span>Mein Konto</span>
            </router-link>
        </div>
        <div v-if="user == true" class="col-sm">
            <a href="#" class="button hidden-sm" @click.prevent="logoutUser()" id="logout">
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
