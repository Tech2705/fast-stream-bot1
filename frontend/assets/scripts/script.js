const linkElm = document.getElementById("link");
const titleElm = document.getElementById("title")
const restElm = document.getElementById("rest")
const metaDataCheckbox = document.getElementById("metadata")
const simpleViewCheckbox = document.getElementById("simpleview")
const settingsBtn = document.getElementById("settingbtn")
const settingsElm = document.getElementById("settings")
const logoElm = document.getElementById("logo")
const video = document.getElementById("video")

document.addEventListener("DOMContentLoaded", function () {
	console.log("hey stop! this is a very important message, only 0000.1% or users can find this message.." +
		"so you are very lucky man. so the message is very secrect!\n" +
		"for security reason i can't say directly here but don't worry, i have an idea\n" +
		"copy this secrect code `61 70 6E 61 20 6B 61 61 6D 20 6B 61 72 6E 61 20 67 61 6E 64 75 21 20 3A 29`\n\n" +
		"open this site https://www.rapidtables.com/convert/number/hex-to-ascii.html\n" +
		"enter the code `61 70 6E 61 20 6B 61 61 6D 20 6B 61 72 6E 61 20 67 61 6E 64 75 21 20 3A 29` in the input field\n" +
		"now click on convert button\n")

	linkElm.innerText = window.location.href
	metaDataCheckbox.addEventListener("change", function () {
		if (this.checked) {
			titleElm.classList.add("hidden")
		} else {
			titleElm.classList.remove("hidden")
		}
	})

	simpleViewCheckbox.addEventListener("change", function () {
		if (this.checked) {
			restElm.classList.add("hidden")
			logoElm.classList.add("hidden")
		} else {
			logoElm.classList.remove("hidden")
			restElm.classList.remove("hidden")
		}
	})

	settingsBtn.addEventListener("click", function () {
		if (settingsElm.classList.contains("-right-[100%]")) {
			settingsElm.classList.remove("-right-[100%]")
			settingsElm.classList.add("right-0")
		} else {
			settingsElm.classList.remove("right-0")
			settingsElm.classList.add("-right-[100%]")
		}
	})

	video.addEventListener("keydown", function (e) {
		if (e.key === "ArrowLeft") {
			video.currentTime -= 10
			e.preventDefault()
		} else if (e.key === "ArrowRight") {
			video.currentTime += 10
			e.preventDefault()
		}

	})
})


function copyLink(link, button) {
	try {
		navigator.clipboard.writeText(link)
		button.innerHTML = `<i class="fa-solid fa-check"></i>`
	} catch (err) {
		console.log("unable to copy , err :", err)
		button.innerHTML = `<i class="fa-solid fa-xmark text-red-400"></i>`
	}

	timeout = setTimeout(() => {
		button.innerHTML = `<i class="fa-solid fa-clipboard"></i>`
		clearTimeout(timeout);
	}, 1000);
}


