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
                                <select v-model="area">
                                    <option v-for="a in areas" :key="a.name">
                                        {{ a.name }}
                                    </option>
                                </select>
                                <br><br>
                                <label for="tType">Tickettyp</label>
                                <select v-model="tType">
                                    <option v-for="tt in tTypes" :key="tt.name">
                                        {{ tt.name }}
                                    </option>
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
            completeUser: {},
            error: [],
            selected: '',
            areas: [
                { id: 1, name: 'SB-Zone-1' },
                { id: 2, name: 'SB-Zone-2' },
                { id: 3, name: 'SB-Zone-3' },
                { id: 4, name: 'B-Zone-1' },
                { id: 5, name: 'B-Zone-2' },
                { id: 6, name: 'B-Zone-3' },
                { id: 7, name: 'B-Zone-4' },
            ],
            tTypes: [
                { id: 1, name: 'Tagesticket' },
                { id: 2, name: 'Monatsticket' },
            ],
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
                            this.validateUser(idToken);
                        }).catch((error) => {
                            console.log(error);
                        });
    },
    methods: {
        buy() {
            let temptType;
            if (this.tType === 'Tagesticket') {
                temptType = 0;
            } else {
                temptType = 1;
            }
            axios.post(`${process.env.VUE_APP_BACKEND_HOST}/buy`, {
                UId: this.completeUser.uid,
                AreaType: this.area,
                TicketType: temptType,
                })
                .then((response) => {
                    console.log(response);
                })
                .catch((e) => {
                    console.log(e);
                    this.error.push(e);
                });
        },
        validateUser(idToken) {
            axios.post(`${process.env.VUE_APP_BACKEND_HOST}/user`, {
                Token: idToken,
                })
                .then((response) => { this.completeUser = response.data; })
                .catch((e) => {
                this.error.push(e);
                });
        },
    },
};
</script>
