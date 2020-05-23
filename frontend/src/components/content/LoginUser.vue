<template>
    <div class="contanier">
        <form>
            <fieldset>
                <legend>Login</legend>
                <input type="text" id="email" placeholder="Email-Adresse" v-model="email" />
                <br />
                <input type="text" id="password" placeholder="Password" v-model="password" />
                <br />
                <button class="primary" @click.prevent="loginUser()">Login</button>
                <button class="primary" @click.prevent="logoutUser()">Logout</button>
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
    </div>
</template>

<script>
import firebase from 'firebase';

const emailRE = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

export default {
    name: 'login',
    data() {
        return {
            email: '',
            password: '',
            error: [],
        };
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
                            // Token zu Bürgerbüro senden -> Uid zurückbekommen -> Dann User valid
                            alert(firebase.auth().currentUser.email);
                        }).catch((error) => {
                            console.log(error);
                        });
                }, (error) => {
                    if (error.code === 'auth/invalid-email'
                    || error.code === 'auth/wrong-password'
                    || error.code === 'auth/user-not-found') {
                        alert('E-Mail oder Passwort falsch oder User existiert nicht');
                    } else if (error.code === 'auth/user-disabled') {
                        alert('Dieser Nutzer ist deaktiviert');
                    } else {
                        alert(error);
                    }
                });
            } else {
                alert('Bitte Mail und Passwort eingeben');
            }
        },
        logoutUser() {
            firebase.auth().signOut().then(() => {
                // Logout erfolgreich
                alert('Logout Erfoglreich!');
            }, (error) => {
                alert('Logout fehlgeschlagen');
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
