<template>
    <section id="home">
        <TheHeader />
        <TheLogin />
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheLogin from '../components/TheLogin.vue';

export default {
    name: 'Login',
    components: {
        TheHeader,
        TheLogin,
    },
};
</script>
