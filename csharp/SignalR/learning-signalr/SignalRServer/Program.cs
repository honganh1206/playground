using System.Text.Json;
using System.Text.Json.Serialization;
using Microsoft.AspNetCore.Http.Connections;
using SignalRServer;
using MessagePack;
using SignalRHubs;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllersWithViews();

builder.Services.AddCors(opts => {
    // Allow only GET requests from any domain
    opts.AddPolicy("AllowAnyGet",
        builder => builder.AllowAnyOrigin()
        .WithMethods("GET")
        .AllowAnyHeader());
    // Allow requests from Example domain only
    opts.AddPolicy("AllowExampleDomain",
        builder => builder.WithOrigins("https://example.com")
            .AllowAnyMethod()
            .AllowAnyHeader()
            .AllowCredentials());
});

builder.Services.AddSignalR(hubOptions => {
    // Ping requests to clients at regular intervals to maintain the connection
    hubOptions.KeepAliveInterval = TimeSpan.FromSeconds(10);
    hubOptions.MaximumReceiveMessageSize = 65_536;
    // Determine how long the server should wait for a response from the client
    hubOptions.HandshakeTimeout = TimeSpan.FromSeconds(15);
    // Adjust num of invocatuibs.hub methods to be called ub parallel
    hubOptions.MaximumParallelInvocationsPerClient = 2;
    hubOptions.EnableDetailedErrors = true;
    // Determine the maximum num of items that can be uploaded to a client-to-server stream
    hubOptions.StreamBufferCapacity = 15;

    if (hubOptions?.SupportedProtocols is not null)
    {
        foreach (var protocol in hubOptions.SupportedProtocols)
        {
            Console.WriteLine($"SignalR supports {protocol} protocol");
        }
    }
})
// Fine-tuning json protocol
.AddJsonProtocol(opt => {
    // Define the policy of naming json props
    opt.PayloadSerializerOptions.PropertyNamingPolicy = null;
    // Accept any implementation of JavaScriptEncoder class from System.Text.Encodings.Web namespace
    opt.PayloadSerializerOptions.Encoder = null;
    // Determine whether or not fields are handled during serialization/deserialization
    opt.PayloadSerializerOptions.IncludeFields = false;

    opt.PayloadSerializerOptions.IgnoreReadOnlyFields = false;
    opt.PayloadSerializerOptions.IgnoreReadOnlyProperties = false;
    // Determine the maximum depth of nested json objects (default is 64 levels)
    opt.PayloadSerializerOptions.MaxDepth = 0; // No limit
    // Determine how numeric data in json messages is handled
    opt.PayloadSerializerOptions.NumberHandling = JsonNumberHandling.Strict; // Numbers are only accepted in a standard json format with NO quotes
    // Determine how keys in C# are translated to json
    opt.PayloadSerializerOptions.DictionaryKeyPolicy = null;
    // Control whether or not props with default values are ignored during serialization/deserialization
    opt.PayloadSerializerOptions.DefaultIgnoreCondition = JsonIgnoreCondition.Never; // Always will be serialized/deserialized
    // Comparison between json fields and C# props will NOT be case-sensitive
    opt.PayloadSerializerOptions.PropertyNameCaseInsensitive = false;
    opt.PayloadSerializerOptions.DefaultBufferSize = 32_768;
    // Determine how comments are handled in json
    opt.PayloadSerializerOptions.ReadCommentHandling = JsonCommentHandling.Skip; // Comments are allowed but not read
    opt.PayloadSerializerOptions.ReferenceHandler = null;
    opt.PayloadSerializerOptions.UnknownTypeHandling = JsonUnknownTypeHandling.JsonElement; // A type declared as object is deserialized as json element
    // Allow json to be written in human-readable format with indentations
    opt.PayloadSerializerOptions.WriteIndented = true;

    Console.WriteLine($"Number of default JSON converters: {opt.PayloadSerializerOptions.Converters.Count}");
})
// .AddMessagePackProtocol(opt => {
//     opt.SerializerOptions = MessagePackSerializerOptions.Standard
//         // Either allow or disallow serialization of untrusted message requests
//         .WithSecurity(MessagePackSecurity.UntrustedData)
//         // Apply compression algorithm to
//         .WithCompression(MessagePackCompression.Lz4Block) // Compress the MessagePack sequence to a single block
//         // Allow assembly version of MessagePack to be different from the one the server is using
//         .WithAllowAssemblyVersionMismatch(true)
//         // Allow old MessagePack specs to be used
//         .WithOldSpec()
//         // Assembly version will not be included in the message
//         .WithOmitAssemblyVersion(true);
// })
.AddStackExchangeRedis("127.0.0.1:6379");

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Home/Error");
    // The default HSTS value is 30 days. You may want to change this for production scenarios, see https://aka.ms/aspnetcore-hsts.
    app.UseHsts();
}

app.UseHttpsRedirection();
app.UseStaticFiles();

app.UseRouting();

app.UseCors("AllowAnyGet").UseCors("AllowExampleDomain");

app.UseAuthorization();

app.MapControllerRoute(
    name: "default",
    pattern: "{controller=Home}/{action=Index}/{id?}");

app.MapHub<LearningHub>("/learningHub", opt => {
    // Allow the availability of SignalR transport mechanisms
    opt.Transports = HttpTransportType.WebSockets | HttpTransportType.LongPolling;
    opt.CloseOnAuthenticationExpiration = true;
    // Max buffer size for data exchange in application layer
    opt.ApplicationMaxBufferSize = 65_536;
    // Max buffer size for data exchange in transport layer
    opt.TransportMaxBufferSize = 65_536;
    // Minimum protocol version supported
    opt.MinimumProtocolVersion = 0; // Any version is supported
    // How long the application will wait for the send action to complete
    opt.TransportSendTimeout = TimeSpan.FromSeconds(10);
    opt.WebSockets.CloseTimeout = TimeSpan.FromSeconds(3);
    opt.LongPolling.PollTimeout = TimeSpan.FromSeconds(10);

    Console.WriteLine($"Authorization data items: {opt.AuthorizationData.Count}");
});
app.UseBlazorFrameworkFiles();
app.Run();
