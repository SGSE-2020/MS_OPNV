<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <!-- The Modal -->
                    <div id="myModal" class="modal">
                        <!-- The Close Button -->
                        <img class="modal-content" id="img01">
                        <div id="caption">
                            <p id="text" style="color: white;"></p>
                            <p id="text2" style="color: white;"></p>
                        </div>
                        <button @click.prevent="buyUser()">Kauf bestätigen</button>
                        <button @click.prevent="hideModal()">Abbrechen</button>
                    </div>
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
                                <button class="primary" @click.prevent="showModal()">Kaufen</button>
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
                { id: 1, name: 'SB-Zone-1', price: 2.5 },
                { id: 2, name: 'SB-Zone-2', price: 1.5 },
                { id: 3, name: 'SB-Zone-3', price: 2.0 },
                { id: 4, name: 'B-Zone-1', price: 1.5 },
                { id: 5, name: 'B-Zone-2', price: 2.5 },
                { id: 6, name: 'B-Zone-3', price: 2.0 },
                { id: 7, name: 'B-Zone-4', price: 3.5 },
            ],
            tTypes: [
                { id: 1, name: 'Tagesticket', code: 0 },
                { id: 2, name: 'Monatsticket', code: 1 },
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
    },
    methods: {
        buyUser() {
            firebase.auth().currentUser.getIdToken(true).then((idToken) => {
                            this.buy(idToken);
                        }).catch((error) => {
                            console.log(error);
                        });
        },
        buy(idToken) {
            let temptType;
            if (this.tType === 'Tagesticket') {
                temptType = 0;
            } else {
                temptType = 1;
            }
            axios.post(`${process.env.VUE_APP_BACKEND_HOST}/buy`, {
                Token: idToken,
                AreaType: this.area,
                TicketType: temptType,
                })
                .then((response) => {
                    console.log(response);
                })
                .catch((e) => {
                    this.error.push(e);
                });
            this.hideModal();
        },
        showModal() {
            let text = document.getElementById('text');
            let text2 = document.getElementById('text2');

            let modal = document.getElementById('myModal');
            modal.style.display = 'block';

            let tempArea = this.areas.find((el) => el.name === this.area);
            let temptType = this.tTypes.find((el) => el.name === this.tType);
            let bill;
            if (temptType.name === 'Tagesticket') {
                bill = tempArea.price;
            } else {
                bill = (tempArea.price * 30) - ((tempArea.price) * 30 * 0.1);
            }
            text.innerHTML = `Sind sie Sicher das Sie ein Tagesticket für ${this.area} kaufen möchten?`;
            text2.innerHTML = `Dafür wird Ihr Konto mit ${bill} € belastet!`;
        },
        hideModal() {
            let modal = document.getElementById('myModal');
            modal.style.display = 'none';
        },
    },
};
</script>
