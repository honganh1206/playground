using System.Diagnostics;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.SignalR;
using SignalRHubs;
using SignalRServer.Models;

namespace SignalRServer.Controllers;

public class HomeController : Controller
{
    private readonly ILogger<HomeController> _logger;

    // Strongly-typed hub context
    private readonly IHubContext<LearningHub, ILearningHubClient> _hubContext;

    public HomeController(ILogger<HomeController> logger, IHubContext<LearningHub, ILearningHubClient> hubContext)
    {
        _logger = logger;
        _hubContext = hubContext;
    }

    public async Task<IActionResult> Index()
    {
        await _hubContext.Clients.All.ReceiveMessage("Index page has been opened by a client.");
        return View();
    }

    public IActionResult Privacy()
    {
        return View();
    }

    [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
    public IActionResult Error()
    {
        return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
    }

    public IActionResult WebAssemblyClient()
    {
        return View();
    }
}
