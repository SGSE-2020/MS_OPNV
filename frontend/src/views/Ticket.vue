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
                                <label for="vType">Verkehrsmittel</label>
                                <select name="vType" id="vType">
                                    <option value="0">Bus</option>
                                    <option value="1">Bahn</option>
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
            vType: '',
            tType: '',
            area: '',
            user: '',
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
        buy() {

        },
        // testBackendCall() {
        //     axios.get(`${process.env.VUE_APP_BACKEND_HOST}/api`)
        //         .then((response) => {
        //             console.log(response.data);
        //             this.text = response.data;
        //         })
        //         .catch((e) => {
        //         this.errors.push(e);
        //         });
        // },
        // testGetUsersCall() {
        //     axios.get(`${process.env.VUE_APP_BACKEND_HOST}/users`)
        //         .then((response) => {
        //             this.users = response.data;
        //         })
        //         .catch((e) => {
        //             this.errors.push(e);
        //         });
        // },
    },
};
</script>
