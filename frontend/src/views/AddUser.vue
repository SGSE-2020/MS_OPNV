<template>
    <section id="home">
        <TheHeader />
        <AddUser />
        <TheFooter />
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheFooter from '../components/TheFooter.vue';
import AddUser from '../components/content/AddUser.vue';

export default {
    components: {
        TheHeader,
        AddUser,
        TheFooter,
    },
};
</script>
