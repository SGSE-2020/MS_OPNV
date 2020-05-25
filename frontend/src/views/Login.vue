<template>
    <section id="home">
        <TheHeader />
        <LoginUser />
        <TheFooter />
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheFooter from '../components/TheFooter.vue';
import LoginUser from '../components/content/LoginUser.vue';

export default {
    components: {
        TheHeader,
        LoginUser,
        TheFooter,
    },
};
</script>
