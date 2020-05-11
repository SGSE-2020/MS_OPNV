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
import AddUser from '../components/AddUser.vue';

export default {
    components: {
        TheHeader,
        TheFooter,
        AddUser,
    },
};
</script>
