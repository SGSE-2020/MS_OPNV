<template>
    <section id="home">
        <TheHeader />
        <Welcome />
        <TheFooter />
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheFooter from '../components/TheFooter.vue';
import Welcome from '../components/content/Welcome.vue';

export default {
    name: 'home',
    components: {
        TheHeader,
        Welcome,
        TheFooter,
    },
};
</script>
