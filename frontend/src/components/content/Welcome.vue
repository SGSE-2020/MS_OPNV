<template>
    <div class="row">
        <div class="col-md-10 drawer" id="welcome">
            <h1>Willkommen auf dem ÖPNV-Portal unserer Smart City</h1>
        </div>
        <TheSidebar />
    </div>
</template>

<script>
import TheSidebar from '../TheSidebar.vue';

export default {
    components: {
        TheSidebar,
    },
};
</script>
