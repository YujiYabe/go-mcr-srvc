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

					'/web/css/index.css',

					'/web/css/externalpackage/bootstrap-vue.css.map',
					'/web/css/externalpackage/bootstrap-vue.css',
					'/web/css/externalpackage/bootstrap.min.css.map',
					'/web/css/externalpackage/bootstrap.min.css',
					'/web/css/externalpackage/fontawesome-all.css',


					'/web/css/webfonts/fa-solid-900.woff2',
					'/web/css/webfonts/fa-solid-900.woff',
					'/web/css/webfonts/fa-solid-900.ttf',
					'/web/css/webfonts/fa-solid-900.svg',
					'/web/css/webfonts/fa-solid-900.eot',
					'/web/css/webfonts/fa-regular-400.woff2',
					'/web/css/webfonts/fa-regular-400.woff',
					'/web/css/webfonts/fa-regular-400.ttf',
					'/web/css/webfonts/fa-regular-400.svg',
					'/web/css/webfonts/fa-regular-400.eot',
					'/web/css/webfonts/fa-brands-400.woff2',
					'/web/css/webfonts/fa-brands-400.woff',
					'/web/css/webfonts/fa-brands-400.ttf',
					'/web/css/webfonts/fa-brands-400.svg',
					'/web/css/webfonts/fa-brands-400.eot',

					'/web/js/index.js',
					'/web/js/externalpackage/axios.min.js',
					'/web/js/externalpackage/axios.min.map',
					'/web/js/externalpackage/bootstrap-vue.js',
					'/web/js/externalpackage/httpVueLoader.js',
					'/web/js/externalpackage/bootstrap-vue.js.map',
					'/web/js/externalpackage/vue-awesome.js',
					'/web/js/externalpackage/vue.js',
					'/web/js/externalpackage/vue.min.js',

					'/web/image/tv/tv_tokyo.png',
					'/web/image/tv/tv_asahi.png',
					'/web/image/tv/tokyo_mx.jpeg',
					'/web/image/tv/tbs.jpeg',
					'/web/image/tv/nihon_tv.png',
					'/web/image/tv/nhk_g.png',
					'/web/image/tv/nhk_e.png',
					'/web/image/tv/fuji_tv.png',

					'/web/vue/home.vue',

					'/web/vue/room/bed.vue',
					'/web/vue/room/living.vue',

					'/web/vue/device/monitor.vue',
					'/web/vue/device/tv.vue',
					'/web/vue/device/raspi.vue',
					'/web/vue/device/ambient.vue',
					'/web/vue/device/projector.vue',

					'/web/vue/icon/circle-down.vue',
					'/web/vue/icon/chevron-up.vue',
					'/web/vue/icon/chevron-down.vue',
					'/web/vue/icon/circle-left.vue',
					'/web/vue/icon/common-arrow-cycle.vue',
					'/web/vue/icon/circle-up.vue',
					'/web/vue/icon/circle-right.vue',
					'/web/vue/icon/common-play.vue',
					'/web/vue/icon/common-pause.vue',
					'/web/vue/icon/common-loading.vue',
					'/web/vue/icon/common-hdd.vue',
					'/web/vue/icon/common-circle.vue',
					'/web/vue/icon/common-reloading.vue',
					'/web/vue/icon/common-powerOnOff.vue',
					'/web/vue/icon/common-return.vue',
					'/web/vue/icon/device-monitor.vue',
					'/web/vue/icon/device-light.vue',
					'/web/vue/icon/device-ambient.vue',
					'/web/vue/icon/device-aircon.vue',
					'/web/vue/icon/common-stop.vue',
					'/web/vue/icon/light-full.vue',
					'/web/vue/icon/light-eco.vue',
					'/web/vue/icon/device-tv.vue',
					'/web/vue/icon/device-raspi.vue',
					'/web/vue/icon/device-projector.vue',
					'/web/vue/icon/media-record.vue',
					'/web/vue/icon/media-cd.vue',
					'/web/vue/icon/light-night.vue',
					'/web/vue/icon/media-replay.vue',
					'/web/vue/icon/menu-main.vue',
					'/web/vue/icon/memu-sub.vue',
					'/web/vue/icon/temperature-cool.vue',
					'/web/vue/icon/room-living.vue',
					'/web/vue/icon/room-bed.vue',
					'/web/vue/icon/operation-forward.vue',
					'/web/vue/icon/operation-f-forward.vue',
					'/web/vue/icon/operation-f-backward.vue',
					'/web/vue/icon/operation-backward.vue',
					'/web/vue/icon/volume-up.vue',
					'/web/vue/icon/volume-down.vue',
					'/web/vue/icon/temperature-warm.vue',
					'/web/vue/icon/common-not-unique.vue',

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
