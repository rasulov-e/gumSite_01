// const productsDOM = document.querySelector('.products')
// const cartItemsDOM = document.querySelector('.cart-items')
// const cartShowBtn = document.querySelector('.cart-icon')
// const cartCloseBtn = document.querySelector('.close-cart')
// const cartOverlay = document.querySelector('.cart-overlay')
// const cartDOM = document.querySelector('.cart')

// let products = []
// let cart = []
// let buttonsDOM = [];

// cartShowBtn.addEventListener("click", function () {
// 	cartOverlay.className = 'cart-overlay-show'

// 	cartDOM.className = 'cart-show'
// })
// cartCloseBtn.addEventListener("click", function () {
// 	cartOverlay.className = 'cart-overlay'

// 	cartDOM.className = 'cart'
// })
// function getBagButtons() {
// 	const buttons = [...document.querySelectorAll(".bag-btn")];
// 	buttonsDOM = buttons;
// 	buttons.forEach(button => {
// 		let id = button.dataset.id;
// 		let inCart = cart.find(item => item.id === id);
// 		if (inCart) {
// 			button.innerHTML = "В корзине";
// 			button.disabled = true;
// 			return
// 		}
// 		button.addEventListener("click", (event) => {
// 			console.log("hello")
// 			event.target.innerText = "В корзине";
// 			event.target.disabled = true;

// 			cart = [...cart, products[id]];
// 		});
// 	});
// }


// function getProductByName(thename) {
// 	fetch('http://localhost:8080/selectProdcutByName', {
// 		method: "GET",
// 		headers: {
// 			"Content-Type": "application/json"
// 		},
// 		body: JSON.stringify({
// 			name: thename
// 		}),
// 	}).then(res => res.json())
// 		.then(data => {
// 			let elements = `
// 			<div class="cart-item">
// 				<img src="./images/${data.name}.png" alt="" style="height: 100%;">
// 				<div>
// 					<h4>${data.name}</h4>
// 					<h5>${data.cost} сом</h5>
// 					<span class="remove-item">убрать</span>
// 				</div>
// 				<div>
// 					<i class="fas fa-chevron-up"></i>
// 					<p class="item-amount">1</p>
// 					<i class="fas fa-chevron-down"></i>
// 				</div>
// 			</div>
// 				`
// 			cartItemsDOM += elements;
// 		});
// }

// function getAllProducts() {
// 	fetch('http://localhost:8080/selectAllProducts', {
// 		method: "GET",
// 	}).then(res => res.json())
// 		.then(data => {
// 			let elements = ""
// 			let i = 0
// 			data.forEach(element => {
// 				products += data;
// 				elements += `
// 				<div class="product">
// 					<img src="./images/${element.name}.png" al="Food" style="height: 200px;">
// 					<div class="product_description">
// 						<h1>${element.name}</h1>
// 						<p class="price">${element.cost} сом</p>
// 						<button class="bag-btn" data=id="${i}">Добавить в корзину</button>
// 					</div>

// 				</div>
// 				`
// 				i++;
// 			});
// 			productsDOM.innerHTML = elements;
// 		})
// }

// document.addEventListener("DOMContentLoaded", () => {

// 	getAllProducts().then(function () {
// 		getBagButtons();
// 	});


// });

const cartBtn = document.querySelector(".cart-btn");
const cartShowBtn = document.querySelector(".cart-icon");
const cartCloseBtn = document.querySelector(".close-cart");
const clearCartBtn = document.querySelector(".clear-cart");
const cartDOM = document.querySelector(".cart");
const cartOverlay = document.querySelector(".cart-overlay");
const cartItems = document.querySelector(".cart-items");
const cartTotal = document.querySelector(".cart-total");
const cartContent = document.querySelector(".cart-content");
const productsDOM = document.querySelector(".products");

let cart = [];
let buttonsDOM = [];
let price = 0;
let products;

class Products {
	async getProducts() {
		try {
			let result = await fetch('http://localhost:8080/selectAllProducts', {
				method: "GET",
			});
			let data = await result.json();
			products = data;
			return products;
		} catch (error) {
			console.log(error);
		}
	}
}

class UI {
	displayProducts(products) {
		let result = "";
		products.forEach(product => {
			result += `
			<div class="product">
 					<img src="./images/${product.name}.png" al="Food" style="height: 200px;">
 					<div class="product_description">
 						<h1>${product.name}</h1>
						<p class="price">${product.cost} сом</p>
					<button class="bag-btn" data-id="${product.id}">Добавить в корзину</button>
					</div>

			</div>
			`
		});
		productsDOM.innerHTML = result;
	}

	getBagButtons() {
		const buttons = [...document.querySelectorAll(".bag-btn")];
		buttonsDOM = buttons;
		buttons.forEach(button => {
			let id = button.dataset.id;
			let inCart = cart.find(item => item.id === id);
			if (inCart) {
				button.innerHTML = "В корзине";
				button.disabled = true;
				return
			}
			button.addEventListener("click", (event) => {
				event.target.innerText = "В корзине";
				event.target.disabled = true;

				let cartItem = { ...Storage.getProduct(id), amount: 1 };

				cart = [...cart, cartItem];

				Storage.saveCart(cart);
				this.setCartValues(cart);
			});
		});
	}
	setCartValues(cart) {
		let tempTotal = 0;
		let itemsTotal = 0;

		cart.forEach(item => {
			tempTotal += Storage.getProduct(item.id);

			itemsTotal++;
		});
		cartTotal.innerText = itemsTotal;
		cartItems.innerText = itemsTotal;
		console.log(cartTotal, cartItems);
	}
}



class Storage {
	static saveProducts(products) {
		localStorage.setItem("products", JSON.stringify(products));
	}
	static getProduct(id) {
		let products = JSON.parse(localStorage.getItem('products'));
		return products.find(product => product.id === id);
	}
	static saveCart() {
		localStorage.setItem("cart", JSON.stringify(cart));
	}
	static getProductAt(index) {
		let products = JSON.parse(localStorage.getItem('products'));

		return products[index];
	}
}

document.addEventListener("DOMContentLoaded", () => {
	const ui = new UI();
	const products = new Products();

	products.getProducts().then(products => {
		ui.displayProducts(products)
		Storage.saveProducts(products);
	}).then(() => {
		ui.getBagButtons();
	});
});


cartShowBtn.addEventListener("click", function () {
	cartOverlay.classList.add('cart-overlay-show')

	cartDOM.classList.add('cart-show')
});
cartCloseBtn.addEventListener("click", function () {
	cartOverlay.classList.remove('cart-overlay')
	cartDOM.classList.remove('cart-show')

});