<script setup lang="ts">
/*Call Components*/
import SalesOverview from '@/components/dashboard/SalesOverview.vue';
import YearlyBreakup from '@/components/dashboard/YearlyBreakup.vue';
import MonthlyEarning from '@/components/dashboard/MonthlyEarnings.vue';
import RecentTransaction from '@/components/dashboard/RecentTransaction.vue';
import ProductPerformance from '@/components/dashboard/ProductPerformance.vue';
import ProductCards from '@/components/dashboard/ProductCards.vue';

definePageMeta({
    middleware: 'auth'
})

// destruct data from useAuth
const { data } = useAuth();
const name = data.value.data.name;
const username = data.value.data.username;
</script>
<template>
    <v-row>
        <v-col cols="12">
            <h2>Welcome, {{ name }}!</h2>
        </v-col>
    </v-row>
</template>