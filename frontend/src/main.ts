import './style.css'
import App from './App.svelte'
import { addMessages, init } from "svelte-i18n";
import en from './locales/en.json';
import es from './locales/es.json';

addMessages('en', en);
addMessages('es', es);

init({
  fallbackLocale: 'en',
  initialLocale: 'en',
});

const app = new App({
  target: document.getElementById('app')
})



export default app

/* REFERENCES:
https://stackoverflow.com/questions/66145620/svelte-i18n-with-snowpack-shows-blank-page-without-errors
*/
