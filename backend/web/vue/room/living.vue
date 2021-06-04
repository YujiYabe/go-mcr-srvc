<template>
	<b-container fluid class="text-light text-center">
		<b-row>
			<b-col
				cols="3"
				:class="[commonClass, methodClassSelected('ambient')]"
				@click="methodChangeDevice('ambient')"
			>
				<icon-device-ambient />
			</b-col>

			<b-col cols="3"
				:class="[commonClass, methodClassSelected('tv')]"
				@click="methodChangeDevice('tv')"
			>
				<icon-device-tv />
			</b-col>

			<b-col
				cols="3"
				:class="[commonClass, methodClassSelected('monitor')]"
				@click="methodChangeDevice('monitor')"
			>
				<icon-device-monitor />
			</b-col>

			<b-col
				cols="3"
				:class="[commonClass, methodClassSelected('Raspi')]"
				@click="methodChangeDevice('Raspi')"
			>
				<icon-device-raspi />
			</b-col>
		</b-row>

		<br />


		<div v-show="device == 'ambient'">
			<device-ambient
				:common-class="commonClass"
				:light-status="lightStatus"
				:aircon-status="airconStatus"
				:aircon-warm="airconWarm"
				:aircon-cool="airconCool"
				@send-command="methodSendCommand"
			/>
		</div>

		<div v-show="device == 'monitor'">
			<device-monitor :common-class="commonClass" @send-command="methodSendCommand" />
		</div>

		<div v-show="device == 'tv'">
			<device-tv :common-class="commonClass" @send-command="methodSendCommand" />
		</div>

		<div v-show="device == 'Raspi'">
			<device-raspi
				:common-class="commonClass"
				:files="files"
				:vlc-progress="vlcProgress"
				:is-downloading="isDownloading"
				:is-uploading="isUploading"
				@send-command="methodSendCommand"
				@file-upload="methodFileUpload"
			/>
		</div>
	</b-container>
</template>

<script>
module.exports = {
	components: {
		'icon-device-light': httpVueLoader('public/vue/icon/device-light.vue'),
		'icon-device-aircon': httpVueLoader('public/vue/icon/device-aircon.vue'),
		'icon-device-ambient': httpVueLoader('public/vue/icon/device-ambient.vue'),
		'icon-device-monitor': httpVueLoader('public/vue/icon/device-monitor.vue'),
		'icon-device-tv': httpVueLoader('public/vue/icon/device-tv.vue'),
		'icon-device-raspi': httpVueLoader('public/vue/icon/device-raspi.vue'),

		'device-light': httpVueLoader('public/vue/device/light.vue'),
		'device-aircon': httpVueLoader('public/vue/device/aircon.vue'),
		'device-ambient': httpVueLoader('public/vue/device/ambient.vue'),
		'device-monitor': httpVueLoader('public/vue/device/monitor.vue'),
		'device-tv': httpVueLoader('public/vue/device/tv.vue'),
		'device-raspi': httpVueLoader('public/vue/device/raspi.vue'),
	},

	props: {
		'commonClass': {type: String, default: ''},
		'device': {type: String, default: ''},
		'lightStatus': {type: String, default: ''},
		'airconStatus': {type: String, default: ''},
		'airconWarm': {type: Number , default: 0},
		'airconCool': {type: Number , default: 0},

		'files': {type: Array, default: []},
		'vlcProgress': {type: Number, default: 0},
		'isDownloading': { type: Boolean, default: false },
		'isUploading': { type: Boolean, default: false },
	},


	methods: {
		methodClassSelected (value) {
			return value == this.device ? 'btn-primary' : 'btn-secondary';
		},

		methodChangeDevice: function(value) {
			this.$emit('change-device', value);
		},

		methodSendCommand: function(value) {
			this.$emit('send-command', value);
		},

		methodFileUpload: function(value) {
			this.$emit('file-upload', value);
		},
	},

}


