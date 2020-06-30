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
                                src="../assets/Bus.jpg"
                                style="width: 300px"
                                alt="Busfahrplan der Smart City"
                                @click.prevent="showModal('busImg')"
                            />
                            <figcaption>Busfahrplan der Smart City</figcaption>
                        </figure>
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
