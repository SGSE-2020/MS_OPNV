<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <h1>Freie Parkplätze der SmartCity</h1>
                    <div>
                        <table>
                        <tr>
                            <th>Name</th>
                            <th>Frei</th>
                            <th>Gesamt</th>
                            <th>Belegt</th>
                        </tr>
                        <tr v-for="space in spaces" :key="space.DisplayName">
                            <td>{{space.DisplayName}}</td>
                            <td class="free">{{space.TotalSpots - space.UtilizedSpots}}</td>
                            <td>{{space.TotalSpots}}</td>
                            <td>{{space.UtilizedSpots}}</td>
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
import axios from 'axios';
import TheHeader from '../components/TheHeader.vue';
import TheSidebar from '../components/TheSidebar.vue';

export default {
    name: 'MyAccount',
    data() {
        return {
            spaces: [],
            error: [],
        };
    },
    created() {
        axios.get(`${process.env.VUE_APP_BACKEND_HOST}/parkspace`)
                            .then((response) => {
                                // console.log(response);
                                this.spaces = response.data;
                            })
                            .catch((e) => {
                                this.error.push(e);
                                console.log('Fehler beim Daten holen');
                                console.log(e);
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
