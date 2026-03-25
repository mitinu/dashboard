// router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import monthlyInformation from "@/page/WorkingCapacityPopulation/MonthlyInformation.vue";
import operationalInformation from "@/page/OperationalInformation/OperationalInformation.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: operationalInformation
    },
    {
        path: '/monthlyInformation',
        name: 'monthlyInformation',
        component: monthlyInformation
    },
    {
        path: '/operationalInformation',
        name: 'operationalInformation',
        component: operationalInformation,
    }
]

const router = createRouter({history: createWebHistory(), routes})
export default router