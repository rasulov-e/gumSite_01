const productsDOM = document.querySelector('.products')
const cartItemsDOM = document.querySelector('.cart-items')
const cartShowBtn = document.querySelector('.cart-icon')
const cartCloseBtn = document.querySelector('.close-cart')
const cartOverlay = document.querySelector('.cart-overlay')
const cartDOM = document.querySelector('.cart')


let cart = []

cartShowBtn.addEventListener("click", function () {
	cartOverlay.className = 'cart-overlay-show'

	cartDOM.className = 'cart-show'
})
cartCloseBtn.addEventListener("click", function () {
	cartOverlay.className = 'cart-overlay'

	cartDOM.className = 'cart'
})
function setupBtns() {
	const buttons = [...document.querySelectorAll(".bag-btn")];
	buttonsDOM = buttons;
	buttons.forEach(button => {
		button.addEventListener('click', () => {
			getProductByName(button.data - id)
			console.log("working")
		})

	});
}

setupBtns()
function getProductByName(thename) {
	fetch('http://localhost:8080/selectProdcutByName', {
		method: "GET",
		headers: {
			"Content-Type": "application/json"
		},
		body: JSON.stringify({
			name: thename
		}),
	}).then(res => res.json())
		.then(data => {
			let elements = `
			<div class="cart-item">
				<img src="./images/${data.name}.png" alt="" style="height: 100%;">
				<div>
					<h4>${data.name}</h4>
					<h5>${data.cost} сом</h5>
					<span class="remove-item">убрать</span>
				</div>
				<div>
					<i class="fas fa-chevron-up"></i>
					<p class="item-amount">1</p>
					<i class="fas fa-chevron-down"></i>
				</div>
			</div>
				`
			cartItemsDOM += elements;
		});
}

function getAllProducts() {
	fetch('http://localhost:8080/selectAllProducts', {
		method: "GET",
	}).then(res => res.json())
		.then(data => {
			let elements = ""
			data.forEach(element => {
				elements += `
				<div class="product">
					<img src="./images/${element.name}.png" al="Food" style="height: 200px;">
					<div class="product_description">
						<h1>${element.name}</h1>
						<p class="price">${element.cost} сом</p>
						<button class="bag-btn" data=id=${element.name}>Добавить в корзину</button>
					</div>
	
				</div>
				`
			});
			productsDOM.innerHTML = elements;
		})
}

getAllProducts()
document.addEventListener("DOMContentLoaded", () => {
	setupBtns()
});