<template>
    <section id="home">
        <TheHeader />
        <div class="row">
            <div class="col-md-10 drawer" id="welcome">
                <h1>Hier können bald Tickets für Bus und Bahn erworben werden!</h1>
            </div>
            <TheSidebar />
        </div>
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheSidebar from '../components/TheSidebar.vue';

export default {
    name: 'Ticket',
    components: {
        TheHeader,
        TheSidebar,
    },
};
</script>
