# 🦠 Virus Game: SEIR Model Simulation

이 프로젝트는 수학적 역학 모델인 SEIR 미분방정식을 기반으로 한 전염병 확산 시뮬레이션 게임입니다. 사용자는 질병의 파라미터를 조절하며 전염병의 확산 양상을 관찰하고 제어할 수 있습니다.

---

## 🧪 Mathematical Background: SEIR Model

본 게임은 감염병의 전파 과정을 다음과 같은 4가지 상태로 구분하여 계산합니다:

* S (Susceptible, 감염 대상군): 아직 감염되지 않았으나 감염될 가능성이 있는 집단
* E (Exposed, 노출군): 바이러스에 노출되어 감염되었으나 아직 전염력은 없는 잠복기 집단
* I (Infectious, 감염군): 전염력을 가지고 병을 퍼뜨리는 집단
* R (Recovered, 회복군): 면역을 획득하거나 완치되어 더 이상 감염되지 않는 집단



시스템의 변화율은 아래의 상미분방정식(ODE)을 통해 계산됩니다:

$$
\begin{aligned}
\frac{dS}{dt} &= -\beta \frac{SI}{N} \\
\frac{dE}{dt} &= \beta \frac{SI}{N} - \sigma E \\
\frac{dI}{dt} &= \sigma E - \gamma I \\
\frac{dR}{dt} &= \gamma I
\end{aligned}
$$

---

## 🎮 Game Features

* Real-time Simulation: SEIR 모델에 따라 실시간으로 변화하는 그래프 및 개체 시뮬레이션.
* Parameter Control: 감염률($\beta$), 잠복기($\sigma$), 회복률($\gamma$) 등을 직접 조절.
* Visualization: 시간에 따른 감염병 확산 추이를 시각적인 그래프로 제공.

---
