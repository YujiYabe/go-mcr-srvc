/*
 *
 *  Air Horner
 *  Copyright 2015 Google Inc. All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License
 *
 */
// importScripts('serviceworker-cache-polyfill.js');

const version = "0.1";
const cacheName = `home-control-${version}`;

self.addEventListener('install', e => {
	const timeStamp = Date.now();
	e.waitUntil(
		caches.open(cacheName).then(cache => {
			return cache.addAll([
					`/`,
					`/index.html`,
					'/manifest.json',

					'/public/css/index.css',

					'/public/css/externalpackage/bootstrap-vue.css.map',
					'/public/css/externalpackage/bootstrap-vue.css',
					'/public/css/externalpackage/bootstrap.min.css.map',
					'/public/css/externalpackage/bootstrap.min.css',
					'/public/css/externalpackage/fontawesome-all.css',


					'/public/css/webfonts/fa-solid-900.woff2',
					'/public/css/webfonts/fa-solid-900.woff',
					'/public/css/webfonts/fa-solid-900.ttf',
					'/public/css/webfonts/fa-solid-900.svg',
					'/public/css/webfonts/fa-solid-900.eot',
					'/public/css/webfonts/fa-regular-400.woff2',
					'/public/css/webfonts/fa-regular-400.woff',
					'/public/css/webfonts/fa-regular-400.ttf',
					'/public/css/webfonts/fa-regular-400.svg',
					'/public/css/webfonts/fa-regular-400.eot',
					'/public/css/webfonts/fa-brands-400.woff2',
					'/public/css/webfonts/fa-brands-400.woff',
					'/public/css/webfonts/fa-brands-400.ttf',
					'/public/css/webfonts/fa-brands-400.svg',
					'/public/css/webfonts/fa-brands-400.eot',

					'/public/js/index.js',
					'/public/js/externalpackage/axios.min.js',
					'/public/js/externalpackage/axios.min.map',
					'/public/js/externalpackage/bootstrap-vue.js',
					'/public/js/externalpackage/httpVueLoader.js',
					'/public/js/externalpackage/bootstrap-vue.js.map',
					'/public/js/externalpackage/vue-awesome.js',
					'/public/js/externalpackage/vue.js',
					'/public/js/externalpackage/vue.min.js',

					'/public/image/tv/tv_tokyo.png',
					'/public/image/tv/tv_asahi.png',
					'/public/image/tv/tokyo_mx.jpeg',
					'/public/image/tv/tbs.jpeg',
					'/public/image/tv/nihon_tv.png',
					'/public/image/tv/nhk_g.png',
					'/public/image/tv/nhk_e.png',
					'/public/image/tv/fuji_tv.png',

					'/public/vue/home.vue',

					'/public/vue/room/bed.vue',
					'/public/vue/room/living.vue',

					'/public/vue/device/monitor.vue',
					'/public/vue/device/tv.vue',
					'/public/vue/device/raspi.vue',
					'/public/vue/device/ambient.vue',
					'/public/vue/device/projector.vue',

					'/public/vue/icon/circle-down.vue',
					'/public/vue/icon/chevron-up.vue',
					'/public/vue/icon/chevron-down.vue',
					'/public/vue/icon/circle-left.vue',
					'/public/vue/icon/common-arrow-cycle.vue',
					'/public/vue/icon/circle-up.vue',
					'/public/vue/icon/circle-right.vue',
					'/public/vue/icon/common-play.vue',
					'/public/vue/icon/common-pause.vue',
					'/public/vue/icon/common-loading.vue',
					'/public/vue/icon/common-hdd.vue',
					'/public/vue/icon/common-circle.vue',
					'/public/vue/icon/common-reloading.vue',
					'/public/vue/icon/common-powerOnOff.vue',
					'/public/vue/icon/common-return.vue',
					'/public/vue/icon/device-monitor.vue',
					'/public/vue/icon/device-light.vue',
					'/public/vue/icon/device-ambient.vue',
					'/public/vue/icon/device-aircon.vue',
					'/public/vue/icon/common-stop.vue',
					'/public/vue/icon/light-full.vue',
					'/public/vue/icon/light-eco.vue',
					'/public/vue/icon/device-tv.vue',
					'/public/vue/icon/device-raspi.vue',
					'/public/vue/icon/device-projector.vue',
					'/public/vue/icon/media-record.vue',
					'/public/vue/icon/media-cd.vue',
					'/public/vue/icon/light-night.vue',
					'/public/vue/icon/media-replay.vue',
					'/public/vue/icon/menu-main.vue',
					'/public/vue/icon/memu-sub.vue',
					'/public/vue/icon/temperature-cool.vue',
					'/public/vue/icon/room-living.vue',
					'/public/vue/icon/room-bed.vue',
					'/public/vue/icon/operation-forward.vue',
					'/public/vue/icon/operation-f-forward.vue',
					'/public/vue/icon/operation-f-backward.vue',
					'/public/vue/icon/operation-backward.vue',
					'/public/vue/icon/volume-up.vue',
					'/public/vue/icon/volume-down.vue',
					'/public/vue/icon/temperature-warm.vue',
					'/public/vue/icon/common-not-unique.vue',

				])
				.then(() => self.skipWaiting());
		})
	);
});



self.addEventListener('activate', (event) => {
	var cacheWhitelist = [cacheName];

	event.waitUntil(
		caches.keys().then((cacheNames) => {
			return Promise.all(
				cacheNames.map((cacheName) => {
					// ホワイトリストにないキャッシュ(古いキャッシュ)は削除する
					if (cacheWhitelist.indexOf(cacheName) === -1) {
						return caches.delete(cacheName);
					}
				})
			);
		})
	);
});

self.addEventListener('fetch', (event) => {
	event.respondWith(
		caches.match(event.request)
		.then((response) => {
			if (response) {
				return response;
			}

			// 重要：リクエストを clone する。リクエストは Stream なので
			// 一度しか処理できない。ここではキャッシュ用、fetch 用と2回
			// 必要なので、リクエストは clone しないといけない
			let fetchRequest = event.request.clone();

			return fetch(fetchRequest)
				.then((response) => {
					if (!response || response.status !== 200 || response.type !== 'basic') {
						return response;
					}

					// 重要：レスポンスを clone する。レスポンスは Stream で
					// ブラウザ用とキャッシュ用の2回必要。なので clone して
					// 2つの Stream があるようにする
					let responseToCache = response.clone();

					caches.open(cacheName)
						.then((cache) => {
							cache.put(event.request, responseToCache);
						});

					return response;
				});
		})
	);
});
