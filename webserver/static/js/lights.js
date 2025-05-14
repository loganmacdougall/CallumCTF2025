const lights = Array.from(document.getElementById("lights").children);

let patterns = null
let i = 0

update_lights = () => {
  lights.forEach((e, ind) => {
    let value = patterns[i][ind]
    if (value === "0") {
      e.classList.remove("light-on")
      e.classList.add("light-off")
    } else {
      e.classList.add("light-on")
      e.classList.remove("light-off")
    }
  })

  i = (i+1) % 10
}

document.body.onload = async () => {
  const res = await fetch("http://localhost:8081/light_pattern")
  const pattern = await res.text()
  patterns = pattern.split("\n")

  setInterval(update_lights, 400)
  update_lights()
}