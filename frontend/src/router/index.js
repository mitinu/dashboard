// router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import monthlyInformation from "@/page/WorkingCapacityPopulation/MonthlyInformation.vue";
import operationalInformation from "@/page/OperationalInformation/OperationalInformation.vue";
import registration from "@/page/Registration.vue";

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
    },
    {
        path: '/registration',
        name: 'registration',
        component: registration,
    }
]

const router = createRouter({history: createWebHistory(), routes})
export default router