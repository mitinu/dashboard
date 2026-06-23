// router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import monthlyInformation from "@/page/WorkingCapacityPopulation/MonthlyInformation.vue";
import operationalInformation from "@/page/OperationalInformation/OperationalInformation.vue";
import registration from "@/page/Registration.vue";
import users from "@/page/Users/Users.vue"
import population from "@/page/Population/Population.vue"
import finance from "@/page/Finance/Finance.vue";

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
    },
    {
        path: '/users',
        name: 'users',
        component: users,
    },
    {
        path: '/population',
        name: 'population',
        component: population,
    },
    {
        path: '/finance',
        name: 'finance',
        component: finance,
    },
]

const router = createRouter({history: createWebHistory(), routes})
export default router