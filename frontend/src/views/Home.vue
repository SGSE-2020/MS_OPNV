<template>
    <section id="home">
        <TheHeader />
        <div class="row">
            <div class="col-md-10 drawer" id="welcome">
                <h1>Willkommen auf dem ÖPNV-Portal unserer Smart City</h1>
            </div>
            <TheSidebar />
        </div>
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheSidebar from '../components/TheSidebar.vue';

export default {
    name: 'Home',
    components: {
        TheHeader,
        TheSidebar,
    },
};
</script>
