using System.Linq;
using System.Security.Claims;
using System.Threading.Tasks;
using AuthProvider.Models;
using IdentityModel;
using IdentityServer4.Extensions;
using IdentityServer4.Models;
using IdentityServer4.Services;
using Microsoft.AspNetCore.Identity;

namespace AuthProvider
{
    /// <summary>
    /// Execute custom logic when an authentication token gets requested
    /// </summary>
    public class UserProfileService : IProfileService
    {
        private readonly IUserClaimsPrincipalFactory<ApplicationUser> _claimsFactory;
        private readonly UserManager<ApplicationUser> _userManager;

        public UserProfileService(IUserClaimsPrincipalFactory<ApplicationUser> claimsFactory, UserManager<ApplicationUser> userManager)
        {
            _claimsFactory = claimsFactory;
            _userManager = userManager;
        }

        public async Task GetProfileDataAsync(ProfileDataRequestContext context)
        {
            var subject = context.Subject.GetSubjectId();

            var user = await _userManager.FindByIdAsync(subject);

            var claimsPrincipal = await _claimsFactory.CreateAsync(user);

            var claimList = claimsPrincipal.Claims.ToList();

            claimList = claimList
                .Where(c => context.RequestedClaimTypes
                    .Contains(c.Type))
                .ToList();

            // Add user-specific claims
            claimList.Add(new Claim(JwtClaimTypes.Email, user.Email));
            claimList.Add(new Claim(JwtClaimTypes.Name, user.UserName));

            if (_userManager.SupportsUserRole)
            {
                foreach (var roleName in await _userManager.GetRolesAsync(user))
                {
                    claimList.Add(new Claim(JwtClaimTypes.Role, roleName));

                    // Add a special claim for admin user
                    if (roleName == "admin")
                    {
                        claimList.Add(new Claim("admin", "true"));
                    }
                }
            }

            context.IssuedClaims = claimList;

        }

        public async Task IsActiveAsync(IsActiveContext context)
        {
            var subject = context.Subject.GetSubjectId();
            var user = await _userManager.FindByIdAsync(subject);

            // Check if user is not null => Active
            context.IsActive = user != null;
        }
    }
}
