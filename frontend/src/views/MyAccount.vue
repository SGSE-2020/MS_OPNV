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
                            <td>{{this.userinfo.gender}}</td>
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
            userinfo: '',
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
                .then((response) => {
                    console.log('User Wurde Validiert');
                })
                .catch((e) => {
                this.error.push(e);
                });
            } else {
                this.user = false;
            }
        });
        let idToken = firebase.auth().currentUser.getIdToken(true);
        axios.post(`${process.env.VUE_APP_BACKEND_HOST}/user`, {
            Token: idToken,
        })
        .then((response) => {
            console.log(response);
            this.userinfo = {
                uid: '6TbzcPavrSNdq1W1qAKqyfhhvxB2',
                gender: 1,
                firstName: 'Max',
                lastName: 'Muster',
                nickName: 'mmuster',
                email: 'exampleuser@test.de',
                birthDate: 'Mon Jan 01 1990 00:00:00 GMT+0000 (Coordinated Universal Time)',
                streetAddress: 'Beispielstraße 12',
                zipCode: '12345',
                city: 'Smart City',
                image: 'data:image/png;base64,iVBORw0',
                isActive: true,
                };
        })
        .catch((e) => {
            this.error.push(e);
        });
    },
    components: {
        TheHeader,
        TheSidebar,
    },
    methods: {
    },

};
</script>
