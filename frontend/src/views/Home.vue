<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <div class="row">
                <div id="center" class="col-sm-8">
                    <h1>Willkommen auf dem ÖPNV-Portal unserer Smart City</h1>
                    <!--<div class="card large">
                    </div>-->
                </div>
                <TheSidebar />
            </div>
        </div>
    </div>
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
