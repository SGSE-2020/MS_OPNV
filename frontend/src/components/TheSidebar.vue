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
                    {{spaces[i - 1].DisplayName}}:
                    {{(spaces[i - 1].TotalSpots) - (spaces[i - 1].UtilizedSpots)}}
                </li>
                <router-link class="button" to="/parkspace">Alle Parkplätze</router-link>
            </ul>
        </div>
        <div class="card small">
            <h3>Werbung</h3>
        </div>
    </div>
</template>


<script>
import axios from 'axios';

export default {
    name: 'TheSidebar',
    data() {
        return {
            spaces: [],
            error: [],
        };
    },
    created() {
        axios.get(`${process.env.VUE_APP_BACKEND_HOST}/fiveparkspaces`)
                            .then((response) => {
                                this.spaces = response.data;
                            })
                            .catch((e) => {
                                this.error.push(e);
                            });
    },
    methods: {
    },

};
</script>
