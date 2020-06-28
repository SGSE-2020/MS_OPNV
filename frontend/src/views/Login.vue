<template>
    <div class="container" >
        <TheHeader />
        <div id="container">
            <TheLogin />
        </div>
    </div>
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
