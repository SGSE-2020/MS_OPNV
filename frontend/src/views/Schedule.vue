<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <h1>Fahrpläne für Bus und Bahn</h1>
                    <div id="container" class="row">
                        <!-- The Modal -->
                        <div id="myModal" class="modal">
                            <!-- The Close Button -->
                            <span class="close" @click="hideModal()">&times;</span>
                            <img class="modal-content" id="img01">
                            <!-- Modal Caption (Image Text) -->
                            <div id="caption"></div>
                        </div>
                        <figure>
                            <img
                                id="bahnImg"
                                src="../assets/Bahn.png"
                                style="width: 300px"
                                alt="Bahnfahrplan der Smart City"
                                @click.prevent="showModal('bahnImg')"
                            />
                            <figcaption>Bahnfahrplan anschauen</figcaption>
                        </figure>
                        <figure>
                            <img
                                id="busImg"
                                src="../assets/Bus.png"
                                style="width: 300px"
                                alt="Busfahrplan der Smart City"
                                @click.prevent="showModal('busImg')"
                            />
                            <figcaption>Busfahrplan der Smart City</figcaption>
                        </figure>
                    </div>
                    <div id="container" class="row">
                        <h3>Übersicht Preis</h3>
                        <p>
                            Bachten Sie: Monatstickets gelten 30 Tage
                            und erhalten 10% Rabatt auf den Tagespreis.
                        </p>
                        <table>
                        <tr>
                            <th>Zone</th>
                            <th>Preis</th>
                            <th>Monatsticket</th>
                        </tr>
                        <tr v-for="area in areas" :key="area.id">
                            <td>{{area.name}}</td>
                            <td>{{area.price}} €</td>
                            <td>{{(area.price * 30) - ((area.price) * 30 * 0.1)}} €</td>
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
import TheHeader from '../components/TheHeader.vue';
import TheSidebar from '../components/TheSidebar.vue';

export default {
    name: 'Schedule',
    data() {
        return {
            areas: [
                { id: 1, name: 'SB-Zone-1', price: 2.5 },
                { id: 2, name: 'SB-Zone-2', price: 1.5 },
                { id: 3, name: 'SB-Zone-3', price: 2.0 },
                { id: 4, name: 'B-Zone-1', price: 1.5 },
                { id: 5, name: 'B-Zone-2', price: 2.5 },
                { id: 6, name: 'B-Zone-3', price: 2.0 },
                { id: 7, name: 'B-Zone-4', price: 3.5 },
            ],
        };
    },
    components: {
        TheHeader,
        TheSidebar,
    },
    methods: {
        showModal(imageId) {
            let img = document.getElementById(imageId);
            let modalImg = document.getElementById('img01');
            let captionText = document.getElementById('caption');

            let modal = document.getElementById('myModal');
            modal.style.display = 'block';
            modalImg.src = img.src;
            captionText.innerHTML = img.alt;
        },
        hideModal() {
            let modal = document.getElementById('myModal');
            modal.style.display = 'none';
        },
    },
};
</script>
