<template>
    <div >
        <form v-if="this.user != true" id="login">
            <fieldset>
                <legend>Login</legend>
                <input type="email" id="email" placeholder="Email-Adresse" v-model="email" />
                <br />
                <input type="password" id="password" placeholder="Password" v-model="password" />
                <br />
                <button class="primary" @click.prevent="loginUser()">Login</button>
                <ul class="errors">
                    <li v-show="!validation.password">
                        Bitte geben Sie ein Passwort ein.
                    </li>
                    <li v-show="!validation.email">
                    Bitte geben Sie eine gültige E-Mail Adresse ein.
                    </li>
                </ul>
            </fieldset>
        </form>
        <div v-if="this.user == true">
            Sie sind schon eingeloggt!
            <button class="primary" @click.prevent="logoutUser()">Logout</button>
        </div>
    </div>
</template>

<script>
import firebase from 'firebase';
import axios from 'axios';

const emailRE = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export default {
    name: 'login',
    data() {
        return {
            email: '',
            password: '',
            user: '',
            error: [],
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
        loginUser() {
            if (
                this.email !== undefined
                && this.email.length > 0
                && this.password !== undefined
                && this.password.length > 0
            ) {
                firebase.auth().signInWithEmailAndPassword(this.email,
                    this.password).then((user) => {
                        firebase.auth().currentUser.getIdToken(true).then((idToken) => {
                            console.log(firebase.auth().currentUser.email);
                            console.log(idToken);
                            this.validateUser(idToken);
                        }).catch((error) => {
                            console.log(error);
                        });
                }, (error) => {
                    if (error.code === 'auth/invalid-email'
                    || error.code === 'auth/wrong-password'
                    || error.code === 'auth/user-not-found') {
                        console.log('E-Mail oder Passwort falsch oder User existiert nicht');
                    } else if (error.code === 'auth/user-disabled') {
                        console.log('Dieser Nutzer ist deaktiviert');
                    } else {
                        console.log(error);
                    }
                });
            } else {
                alert('Bitte Mail und Passwort eingeben');
            }
        },
        logoutUser() {
            firebase.auth().signOut().then(() => {
                // Logout erfolgreich
                console.log('Logout Erfoglreich!');
            }, (error) => {
                console.log('Logout fehlgeschlagen');
            });
        },
        validateUser(idToken) {
            axios.post(`${process.env.VUE_APP_BACKEND_HOST}/user`, {
                Token: idToken,
                })
                .then((response) => { console.log('User Wurde Validiert'); })
                .catch((e) => {
                this.error.push(e);
                });
        },
    },
    computed: {
        validation() {
            return {
                password: !!this.password.trim(),
                email: emailRE.test(this.email),
            };
        },
        isValid() {
            let { validation } = this.validation;
            return Object.keys(validation).every((key) => validation[key]);
        },
    },
};
</script>
