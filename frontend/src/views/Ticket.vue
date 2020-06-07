<template>
    <section id="home">
        <TheHeader />
        <div class="row">
            <div class="col-md-10 drawer" id="welcome">
                <h1>Hier können bald Tickets für Bus und Bahn erworben werden!</h1>
                <button class="primary" @click="testBackendCall()">TestBackendCall</button>
                <p v-if="this.text != ''">{{this.text}}</p>
            </div>
            <TheSidebar />
        </div>
    </section>
</template>

<script>
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
            text: '',
        };
    },
    methods: {
        testBackendCall() {
            axios.get('http://opnv.dvess.network/api/api')
                .then((response) => {
                    console.log(response.data);
                    this.text = response.data;
                })
                .catch((e) => {
                this.errors.push(e);
                });
        },
    },
};
</script>
