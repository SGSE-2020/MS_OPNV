<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <div v-if="this.user == true" id="ticket">
                        <form>
                            <fieldset>
                                <legend>Tickets kaufen</legend>
                                <label for="area">Bereich</label>
                                <select name="area" id="area">
                                    <option value="0">Bereich 1</option>
                                    <option value="1">Bereich 2</option>
                                    <option value="2">Bereich 3</option>
                                    <option value="3">Bereich 4</option>
                                </select>
                                <br><br>
                                <label for="tType">Tickettyp</label>
                                <select name="tType" id="tType">
                                    <option value="0">Tagesticket</option>
                                    <option value="1">Monatsticket</option>
                                </select>
                                <br><br>
                                <button class="primary" @click.prevent="buy()">Kaufen</button>
                            </fieldset>
                        </form>
                    </div>
                    <div v-if="this.user == false">
                        Um Tickets zu kaufen müssen sie sich einloggen!
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
    name: 'Ticket',
    components: {
        TheHeader,
        TheSidebar,
    },
    data() {
        return {
            tType: '',
            area: '',
            user: '',
            userToken: '',
            userId: '',
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
                            this.userToken = idToken;
                        }).catch((error) => {
                            console.log(error);
                        });
        axios.post(`${process.env.VUE_APP_BACKEND_HOST}/user`, {
                Token: this.userToken,
                })
                .then((response) => {
                    this.userId = response.data.userId;
                })
                .catch((e) => {
                    console.log(e);
                this.error.push(e);
                });
    },
    methods: {
        buy() {
            axios.get(`${process.env.VUE_APP_BACKEND_HOST}/buy`, {
                // UId: this.userId,
                UId: '1234',
                AreaType: 'SB-Zone-1',
                TicketType: 0,
                })
                .then((response) => {
                    console.log(response);
                })
                .catch((e) => {
                    console.log(e);
                    this.error.push(e);
                });
        },
    },
};
</script>
