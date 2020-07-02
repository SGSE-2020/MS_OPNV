<template>
    <div class="col-sm-4 col-md">
        <div class="card small">
            <h3>Verkehrsinfo</h3>
            <ul>
                <li>Baustelle - Hansestraße vom 01.06. - 31.06.2020</li>
                <li><a href="#">mehr...</a></li>
            </ul>
        </div>
        <div class="card small">
            <h3>Freie Parkplätze</h3>
            <ul>
                <li v-for="i in 5" :key="i">
                    {{spaces_side[i - 1].DisplayName}}:
                    {{(spaces_side[i - 1].TotalSpots) - (spaces_side[i - 1].UtilizedSpots)}}
                </li>
                <router-link class="button" to="/parkspace">Alle Parkplätze</router-link>
            </ul>
        </div>
        <div class="card small">
            <h3>Werbung</h3>
            <div class="card small">
                <img src="../assets/restaurant_bg.png"/>
                <span style="text-align: center;">Neue Karte seit heute!</span>
                <br/>
                <a
                class="button"
                href="https://restaurants.dvess.network/#/"
                target= "_blank"
                style="text-align: center;"
                >Tisch reservieren</a>
            </div>
            <div class="card small">
                <img src="../assets/supermarkt.png"/>
                <span style="text-align: center;">
                    Sparen Sie heute bis zu 15% auf ihren Einkauf!
                </span>
                <br/>
                <a
                class="button"
                href="https://supermarkt.dvess.network/#"
                target= "_blank"
                style="text-align: center;"
                >Jetzt Einkaufen</a>
            </div>
        </div>
    </div>
</template>


<script>
import axios from 'axios';

export default {
    name: 'TheSidebar',
    data() {
        return {
            spaces_side: [],
            error: [],
        };
    },
    created() {
        axios.get(`${process.env.VUE_APP_BACKEND_HOST}/fiveparkspaces`)
                            .then((response) => {
                                this.spaces_side = response.data;
                            })
                            .catch((e) => {
                                this.error.push(e);
                            });
    },
    methods: {
    },

};
</script>
