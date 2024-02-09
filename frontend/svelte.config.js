import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
export default {
    kit: {
        adapter: adapter({
            // default options are shown. On some platforms
            // these options are set automatically — see below
            pages: '../static',
            assets: '../static',
            fallback: '404.html',
            precompress: false,
            strict: true
        }),
        // paths: {
        //     base: '/static'
        // }
    },
    preprocess: vitePreprocess()
};
