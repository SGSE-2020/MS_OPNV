<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <div v-if="this.user == true" id="ticket">
                        <h1>Mein Konto</h1>
                        <div>
                            <table>
                            <tr>
                                <th></th>
                                <th></th>
                            </tr>
                            <tr>
                                <td>Geschlecht</td>
                                <td v-if="this.userinfo.gender == 0">weiblich</td>
                                <td v-if="this.userinfo.gender == 1">männlich</td>
                            </tr>
                            <tr>
                                <td>Vorname</td>
                                <td>{{this.userinfo.firstName}}</td>
                            </tr>
                            <tr>
                                <td>Nachname</td>
                                <td>{{this.userinfo.lastName}}</td>
                            </tr>
                            <tr>
                                <td>Nickname</td>
                                <td>{{this.userinfo.nickName}}</td>
                            </tr>
                            <tr>
                                <td>Email</td>
                                <td>{{this.userinfo.email}}</td>
                            </tr>
                            <tr>
                                <td>Geburtsdatum</td>
                                <td>{{this.userinfo.birthDate}}</td>
                            </tr>
                            <tr>
                                <td>Adresse</td>
                                <td>
                                    {{this.userinfo.streetAddress}},
                                    {{this.userinfo.zipCode}}
                                    {{this.userinfo.city}}
                                    </td>
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
                                <td>SB-Zone-2</td>
                                <td>Tagesticket</td>
                                <td>03.07.2020</td>
                            </tr>
                            </table>
                        </div>
                        </div>
                    <div v-if="this.user == false">
                        Loggen Sie sich ein um diesen Bereich einsehen zu können!
                    </div>
                </div>
                <!--<TheSidebar />-->
            </div>
        </div>
    </div>
</template>

<script>
import firebase from 'firebase';
import axios from 'axios';
import TheHeader from '../components/TheHeader.vue';
// import TheSidebar from '../components/TheSidebar.vue';

export default {
    name: 'MyAccount',
    data() {
        return {
            user: '',
            userinfo: {},
            tickets: [],
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
        firebase.auth().currentUser.getIdToken(true).then((idToken) => {
                            this.getTickets(idToken);
                            this.getUser(idToken);
                        }).catch((error) => {
                            console.log(error);
                        });
    },
    components: {
        TheHeader,
        // TheSidebar,
    },
    methods: {
        getTickets(idToken) {
            axios.post(`${process.env.VUE_APP_BACKEND_HOST}/ticket`, {
                Token: idToken,
                })
                .then((response) => {
                    console.log(response);
                    this.tickets = response.data;
                })
                .catch((e) => {
                    this.error.push(e);
                });
        },
        getUser(idToken) {
            axios.post(`${process.env.VUE_APP_BACKEND_HOST}/user`, {
                Token: idToken,
                })
                .then((response) => {
                    this.userinfo = response.data;
                })
                .catch((e) => {
                    this.error.push(e);
                });
        },
    },

};
</script>
