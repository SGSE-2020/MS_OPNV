<template>
    <section id="home">
        <TheHeader />
        <div class="row">
            <figure>
                <img src="../assets/Bahn.png" style="width: 500px" alt="Image description"/>
                <figcaption>Bahnfahrplan der Smart City</figcaption>
            </figure>
            <figure>
                <img src="../assets/Bus.jpg" style="width: 500px" alt="Image description"/>
                <figcaption>Busfahrplan der Smart City</figcaption>
            </figure>
            <AddUser />
            <TheSidebar />
        </div>
    </section>
</template>

<script>
import TheHeader from '../components/TheHeader.vue';
import TheSidebar from '../components/TheSidebar.vue';
import AddUser from '../components/AddUser.vue';

export default {
    name: 'Schedule',
    components: {
        TheHeader,
        TheSidebar,
        AddUser,
    },
};
</script>
