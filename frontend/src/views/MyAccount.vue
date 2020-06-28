<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <h1>Mein Konto</h1>
                    <div>
                        <table>
                        <tr>
                            <th></th>
                            <th></th>
                        </tr>
                        <tr>
                            <td>Geschlecht</td>
                            <td>1</td>
                        </tr>
                        <tr>
                            <td>Vorname</td>
                            <td>Max</td>
                        </tr>
                        <tr>
                            <td>Nachname</td>
                            <td>Muster</td>
                        </tr>
                        <tr>
                            <td>Nickname</td>
                            <td>mmuster</td>
                        </tr>
                        <tr>
                            <td>Email</td>
                            <td>exampleuser@test.de</td>
                        </tr>
                        <tr>
                            <td>Geburtsdatum</td>
                            <td>01. Januar 1990</td>
                        </tr>
                        <tr>
                            <td>Adresse</td>
                            <td>Beispielstraße 12, 12345 Smart City</td>
                        </tr>
                        </table>
                    </div>
                    <h1>Meine Tickets</h1>
                    <div>
                        <table style="width:100%">
                        <tr>
                            <th>Bereich</th>
                            <th>Ticketart</th>
                            <th>Gültig bis</th>
                        </tr>
                        <tr>
                            <td>Bahn-4</td>
                            <td>Tagesticket</td>
                            <td>02.07.2020</td>
                        </tr>
                        </table>
                    </div>
                    <h1>Abglaufene Tickets</h1>
                    <div>
                        <table style="width:100%">
                        <tr>
                            <th>Bereich</th>
                            <th>Ticketart</th>
                            <th>Gültig bis</th>
                        </tr>
                        <tr>
                            <td>Bahn-4</td>
                            <td>Tagesticket</td>
                            <td>02.07.2020</td>
                        </tr>
                        </table>
                    </div>
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
    name: 'MyAccount',
    data() {
        return {
            user: '',
            error: [],
        };
    },
    created() {
        firebase.auth().onAuthStateChanged((user) => {
            if (user) {
                this.user = true;
                let idToken = firebase.auth().currentUser.getIdToken(true);
                axios.post(`${process.env.VUE_APP_BACKEND_HOST}/user`, {
                    Token: idToken,
                })
                .then((response) => { console.log('User Wurde Validiert'); })
                .catch((e) => {
                this.error.push(e);
                });
            } else {
                this.user = false;
            }
        });
    },
    components: {
        TheHeader,
        TheSidebar,
    },
};
</script>
